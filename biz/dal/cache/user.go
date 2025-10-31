package cache

import (
	"context"
	"fmt"
	"judgeMore/pkg/errno"
	"math/rand"
	"strings"
	"time"
)

func IsKeyExist(ctx context.Context, key string) bool {
	return ca.Exists(ctx, key).Val() == 1
}

func GetCodeCache(ctx context.Context, key string) (code string, err error) {
	value, err := ca.Get(ctx, key).Result()
	if err != nil {
		return "", errno.NewErrNo(errno.InternalRedisErrorCode, "write code to cache error:"+err.Error())
	}
	var storedCode string
	parts := strings.Split(value, "_")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid code format, expected 2 parts, got %d", len(parts))
	}
	storedCode = parts[0]
	return storedCode, nil
}
func PutCodeToCache(ctx context.Context, key string) (code string, err error) {
	code = generateRandomCode(6)
	timeNow := time.Now().Unix()
	value := fmt.Sprintf("%s_%d", code, timeNow)
	expiration := 10 * time.Minute
	err = ca.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return "", errno.NewErrNo(errno.InternalRedisErrorCode, "write code to cache error:"+err.Error())
	}
	return code, nil
}

// 生成指定位数的随机验证码（字母+数字）
func generateRandomCode(length int) string {
	// 字符集：26个小写字母 + 26个大写字母 + 10个数字
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 初始化随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charSet[r.Intn(len(charSet))]
	}

	return string(code)
}
