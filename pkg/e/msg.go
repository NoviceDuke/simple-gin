package e

var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "fail",
	INVALID_PARAMS:           "請求參數錯誤",
	ERROR_EXIST_CATEGORY:     "已存在該分類",
	ERROR_NOT_EXIST_CATEGORY: "該分類不存在",
	ERROR_NOT_EXIST_PRODUCT:  "該商品不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
