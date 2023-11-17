package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"

	server2 "github.com/RunningShrimp/sim-go/server/internal/route"
	"github.com/bytedance/sonic"
)

// http 处理引擎，不对外暴露
func (s *SimHTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	// 1. 获取请求方法与url
	httpMethod := request.Method
	urlStr := request.URL.Path
	// 2. 根据请求方法和url获取handler
	handler := s.Route.GetRoute(urlStr, httpMethod)
	if handler == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	s.dispatchRequest(writer, request, handler)

}
func (s *SimHTTPServer) dispatchRequest(writer http.ResponseWriter, request *http.Request, handler *server2.Handler) {
	if s.Route == nil {
		panic("请注册路由")
	}
	data := make(map[string]any)
	bodyDataBytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("读取请求体错误：err:%v", err)
		return
	}

	defer func(bodyData io.ReadCloser) {
		err := bodyData.Close()
		if err != nil {
			log.Printf("EasyGoServeHTTP.ServeHTTP error")
		}
	}(request.Body)

	if len(bodyDataBytes) > 0 && getMediaType(request) == ApplicationJSON {
		if err := sonic.Unmarshal(bodyDataBytes, data); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			log.Printf("请求体格式化错误")
		}
	}
	// 3. 获取请求参数
	argValues := make([]reflect.Value, 0)
	// 4. 将请求参数注入到handler参数中
	for _, e := range handler.In {
		fmt.Println(*e)
		argValues = append(argValues, s.dataMapStruct(data, *e))
	}
	// 5. 执行handler
	resultArr := handler.HFunc.Call(argValues)
	// 6. 获取handler执行结果，返回response
	// for _, v := range resultArr {
	//	// TODO.md: 检查error
	//
	//}
	if len(resultArr) > 0 {
		val := resultArr[0]

		switch val.Kind() {
		case reflect.Slice:
			_, _ = fmt.Fprintf(writer, "%v", val.String())
			return
		case reflect.Bool:
			_, _ = fmt.Fprintf(writer, "%v", val.Bool())
			return
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			_, _ = fmt.Fprintf(writer, "%d", val.Int())
			return
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
			_, _ = fmt.Fprintf(writer, "%d", val.Uint())
			return
		case reflect.Float32:
		case reflect.Float64:
			_, _ = fmt.Fprintf(writer, "%f", val.Float())
			return
		case reflect.String:

			_, _ = fmt.Fprintf(writer, "%s", val.String())
			return
		case reflect.Struct:
			bytes, err := json.Marshal(val.Bytes())
			if err != nil {
				return
			}
			_, _ = fmt.Fprintf(writer, "%s", string(bytes))
			return
		default:
			writer.Write(val.Bytes())
			writer.WriteHeader(http.StatusOK)
		}

	} else {
		writer.WriteHeader(http.StatusOK)
	}

}

func (s *SimHTTPServer) dataMapStruct(data map[string]any, argType reflect.Type) reflect.Value {
	val := reflect.New(argType)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < argType.NumField(); i++ {
		t := argType.Field(i)
		f := val.Field(i)
		tag := t.Tag.Get("json")
		if v, ok := data[tag]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			fmt.Println(dataType)
			structType := f.Type()
			fmt.Println(structType)
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					switch structType.Kind() {
					case reflect.Int:
					case reflect.Int8:
					case reflect.Int16:
					case reflect.Int32:
					case reflect.Int64:
						v, err := strconv.ParseInt(v.(string), 10, 64)
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							// TODO.md：未来这里需要优化
							log.Println("数据格式错误")

							break
						}
						f.SetInt(v)
						break
					case reflect.Float32:
					case reflect.Float64:
						v, err := strconv.ParseFloat(v.(string), 64)
						if err != nil {
							log.Println("数据格式错误")

							break
						}

						f.SetFloat(v)
						break
					case reflect.Uint:
					case reflect.Uint8:
					case reflect.Uint16:
					case reflect.Uint32:
					case reflect.Uint64:
						v, err := strconv.ParseUint(v.(string), 10, 64)
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							// TODO.md：未来这里需要优化
							log.Println("数据格式错误")

							break
						}
						f.SetUint(v)
						break
					case reflect.Bool:
						v, err := strconv.ParseBool(v.(string))
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							// TODO.md：未来这里需要优化
							log.Println("数据格式错误")
							break
						}
						f.SetBool(v)
						break
					default:

						panic(t.Name + " type mismatch")
					}
				}
			}
		}
	}
	return val
}
