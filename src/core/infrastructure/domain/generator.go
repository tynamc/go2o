/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package domain

import (
	"fmt"
	"github.com/jrsix/gof/crypto"
	"github.com/jrsix/gof/util"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	minRand int = 100000
	maxRand int = 999900
)

//新订单号
func NewOrderNo(partnerId int) string {
	//PartnerId的首位和末尾再加7位随机数
	unix := time.Now().Unix()
	rand.Seed(unix)
	rd := minRand + rand.Intn(maxRand-minRand) //minRand - maxRand中间的随机数
	timeStr := time.Now().Format("0601")
	ptStr := strconv.Itoa(partnerId)
	return fmt.Sprintf("%s%s%s%d", timeStr, ptStr[:1], ptStr[len(ptStr)-1:], rd)
}

// 创建邀请码(6位)
func GenerateInvitationCode() string {
	var seed string = fmt.Sprintf("%d%s", time.Now().Unix(), util.RandString(6))
	var md5 = crypto.Md5([]byte(seed))
	return md5[8:14]
}

// 创建API编号(10位)
func NewApiId(id int) string {
	var offset = id*360 + id%2
	return fmt.Sprintf("60%s%d", strings.Repeat("0", 8-len(strconv.Itoa(offset))), offset)
}

//创建密钥(16位)
func NewSecret(hex int) string {
	str := fmt.Sprintf("%d$%d", hex, time.Now().Add(time.Hour*24*365).Unix())
	return crypto.Md5([]byte(str))[8:24]
}
