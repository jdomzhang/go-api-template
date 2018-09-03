package util

// import (
// 	"fmt"
// 	"log"

// 	"../../config"
// 	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
// 	"github.com/tobyzxj/uuid"
// )

// // modify it to yours
// var (
// 	ACCESSID        string
// 	ACCESSKEY       string
// 	SIGNNAME        string
// 	TEMPLATE        string
// 	TEMPLATESUCCESS string
// )

// func init() {
// 	ACCESSID = config.All["aliyun.dysms.id"]
// 	ACCESSKEY = config.All["aliyun.dysms.secret"]
// 	SIGNNAME = config.All["aliyun.dysms.signname"]
// 	TEMPLATE = config.All["aliyun.dysms.templatecode"]
// 	TEMPLATESUCCESS = config.All["aliyun.dysms.templatecode.success"]
// }

// // SendSms will send sms by given phone & code
// func SendSms(phone string, code string) error {
// 	if phone == "" {
// 		return fmt.Errorf("phone should not be empty")
// 	}
// 	dysms.HTTPDebugEnable = true
// 	dysms.SetACLClient(ACCESSID, ACCESSKEY)

// 	log.Println("Sending to", phone, "with", code)

// 	respSendSms, err := dysms.SendSms(uuid.New(), phone, SIGNNAME, TEMPLATE, `{"code": "`+code+`"}`).DoActionWithException()
// 	if err != nil {
// 		log.Println("send sms failed", err, respSendSms.Error())
// 		return err
// 	}
// 	log.Println("send sms succeed", respSendSms.GetRequestID())

// 	return nil
// }

// // SendSuccessSms will send sms by given phone & code
// func SendSuccessSms(phone string) error {
// 	if phone == "" {
// 		return fmt.Errorf("phone should not be empty")
// 	}
// 	dysms.HTTPDebugEnable = true
// 	dysms.SetACLClient(ACCESSID, ACCESSKEY)

// 	log.Println("Sending success to", phone)

// 	respSendSms, err := dysms.SendSms(uuid.New(), phone, SIGNNAME, TEMPLATESUCCESS, ``).DoActionWithException()
// 	if err != nil {
// 		log.Println("send sms failed", err, respSendSms.Error())
// 		return err
// 	}
// 	log.Println("send sms succeed", respSendSms.GetRequestID())

// 	return nil
// }
