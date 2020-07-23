package studentManage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"runtime/debug"
	"time"
	logs "yuncode/fabric/logs/yxlogs"
)

/**
 * @Author: WuNaiChi
 * @Date: 2020/6/28 14:35
 * @Desc:
 */
type DefaultController struct {
	beego.Controller
}
type SDKClient struct {
	client    *channel.Client
	ledgerCli *ledger.Client
}

const key = "%s-%s-%s"

// todo:这个是干什么用的
//type ClientCache struct {
//	cache *cache2.Cache
//}

//var Cache *ClientCache

type BCTxDetail struct {
	Txid      string `json:"txId"`
	TxStatus  string `json:"txStatus"` // valid
	Timestamp string `json:"timestamp"`
	BlockNum  int    `json:"blockNum"`
}

func (c *DefaultController) ReturnError(message *Response) {
	c.Ctx.Output.SetStatus(message.Code)
	c.Data["json"] = message
	c.ServeJSON()
}
func (c *DefaultController) ReturnOK(data interface{}) {
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = &Response{
		Code: 200,
		Data: data,
	}
	c.ServeJSON()
}

var expires time.Duration

//func Set(customerID, chainId, orgName string, client *SDKClient) {
//	id := fmt.Sprintf(key, customerID, chainId, orgName)
//	clientCache.cache.Set(id, client, expires)
//}
//
//func (clientCache *ClientCache) Get(customerID, chainId, orgName string) (*SDKClient, bool) {
//	id := fmt.Sprintf(key, customerID, chainId, orgName)
//	client, exist := clientCache.cache.Get(id)
//	if exist {
//		return client.(*SDKClient), true
//	}
//	return nil, false
//}

//func NewService(stuId, chainId, orgName string) (*SDKClient, error) {

//client, exist := Get(stuId, chainId, orgName)
//if !exist {
//	client, err := NewClient(stuId, chainId, orgName)
//	if err != nil {
//		return nil, err
//	}
//	Set(stuId, chainId, orgName, client)
//}
//return client, nil
//}
func NewService(constomerId, chainId, orgName string) (*SDKClient, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("panic stack info:", string(debug.Stack()))
		}
	}()

	if chainId == "" {
		chainId = beego.AppConfig.String("defaultChainID")
	}
	if orgName == "" {
		orgName = beego.AppConfig.String("defaultOrgName")
	}
	configPath := beego.AppConfig.String("sdkConfig")
	if configPath == "" {
		logs.Error("Failed to found the sdk config")
		return nil, errors.New("Failed to found the sdk config")
	}

	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		fmt.Println("newfabsdk err line100", err)
		return nil, err
	}
	channelContext := sdk.ChannelContext(chainId, fabsdk.WithUser(constomerId), fabsdk.WithOrg(orgName))
	ledCli, err := ledger.New(channelContext)
	if err != nil {
		fmt.Println("ChannelContext err line105", err)
		return nil, err
	}
	chainCli, err := channel.New(channelContext)
	if err != nil {
		fmt.Println("ChannelNew err line111", err)
		return nil, err
	}
	return &SDKClient{client: chainCli, ledgerCli: ledCli}, nil
}
func (s *SDKClient) Commit(student *StudentInfo, chaincodeId string) error {
	var err error

	data, err := json.Marshal(student)
	if err != nil {
		return err
	}
	err = s.Invoke(chaincodeId, CommitStudentInfo, []string{string(data)}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *SDKClient) Invoke(chainCodeId, fcn string, args []string, transientMap map[string][]byte) error {

	req := channel.Request{
		ChaincodeID:  chainCodeId,
		Fcn:          fcn,
		Args:         convertArgs2byte(args),
		TransientMap: transientMap,
	}
	_, err := c.client.Execute(req)
	if err != nil {
		return err
	}

	return nil
}
func (c *SDKClient) Query(chainCodeId, fcn string, args []string) ([]byte, error) {
	req := channel.Request{
		ChaincodeID: chainCodeId,
		Fcn:         fcn,
		Args:        convertArgs2byte(args),
	}
	response, err := c.client.Query(req)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (s *SDKClient) Update(student *StudentInfo, chaincodeId string) error {
	data, err := json.Marshal(student)
	if err != nil {
		return err
	}
	err = s.Invoke(chaincodeId, UpdateStudentInfo, []string{string(data)}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SDKClient) QueryList(student *QueryStudentList, chaincodeId string) (*RspQueryStudentInfo, error) {
	data, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}
	rep, err := s.Query(chaincodeId, QueryStudentInfoList, []string{string(data)})
	if err != nil {
		return nil, err
	}
	var studentInfo = new(RspQueryStudentInfo)
	err = json.Unmarshal(rep, studentInfo)
	if err != nil {
		return nil, err
	}
	return studentInfo, nil
}

func (s *SDKClient) QueryInfo(student *QueryStudentInfo, chaincodeId string) (*StudentInfo, error) {
	data, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}
	rep, err := s.Query(chaincodeId, QueryStuInfo, []string{string(data)})
	if err != nil {
		return nil, err
	}
	var studentInfo = new(StudentInfo)
	err = json.Unmarshal(rep, studentInfo)
	if err != nil {
		return nil, err
	}
	return studentInfo, nil
}

func (s *SDKClient) DeleteInfo(student *DeleteStudentInfo, chaincodeId string) error {
	data, err := json.Marshal(student)
	if err != nil {
		return err
	}
	err = s.Invoke(chaincodeId, DeleteStuInfo, []string{string(data)}, nil)
	if err != nil {
		return err
	}
	return nil
}
func convertArgs2byte(str []string) [][]byte {
	var args [][]byte
	for i := 0; i < len(str); i++ {
		args = append(args, []byte(str[i]))
	}
	return args
}
