package utils

import (
	"github.com/google/uuid"
	"strings"
)

const EMPTY_STR string = ""
const HYPHEN_STR string = "-"

/**
 * 获取UUID字符串
 *
 * @return
 */
func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), HYPHEN_STR, EMPTY_STR)
}
