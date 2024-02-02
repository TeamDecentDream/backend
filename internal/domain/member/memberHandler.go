package member

import (
	"backend/internal/models"
	"backend/internal/utils/jwt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetAllMembersHandler(c *gin.Context) {

}

func GetMemberDetailHandler(c *gin.Context) {

}

func LoginHandler(c *gin.Context) {
	var credential models.Credential
	err := c.BindJSON(&credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	kakaoHandler(&credential, c)
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteMemberHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	err = DeleteMember(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func ConnectAddressHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _, _, _, err := jwt.AccessTokenVerifier(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	var input map[string]interface{}
	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}

	addr := input["address"].(string)
	token, err = SaveAddress(id, addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"addrToken": token})
}

func kakaoHandler(credential *models.Credential, c *gin.Context) {
	log.Println(credential)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", "b1408601031a21452f37b4ad7e4009db")
	data.Set("redirect_uri", "http://localhost:3000/login")
	data.Set("code", credential.Code)
	data.Set("client_secret", "bb9bCee45cyOcJsvbr7bfdIcs1AoAC8v")

	resp, err := http.PostForm("https://kauth.kakao.com/oauth/token", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	var responseData models.KakaoOAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "KakaoOAuthResponse JSON parsing Error"})
	}

	log.Println("Kakao OAuth 응답:", responseData)

	userInfoUrl := "https://kapi.kakao.com/v2/user/me"
	req, err := http.NewRequest("GET", userInfoUrl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", "Bearer "+responseData.AccessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	userInfoResp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo models.KakaoMemberResponse
	err = json.NewDecoder(userInfoResp.Body).Decode(&userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 정보를 JSON으로 파싱하는 중 오류 발생"})
		return
	}

	log.Println("Kakao Member 응답:", userInfo.Properties.Nickname)
	log.Println("Kakao Member 응답:", userInfo.KakaoAccount.Email)

	result, err := FindByNameAndEmail(userInfo.Properties.Nickname, userInfo.KakaoAccount.Email)
	if err != nil {
		if strings.Contains(err.Error(), "member not found") {
			var newMember models.Member
			newMember.Name = userInfo.Properties.Nickname
			newMember.Email = userInfo.KakaoAccount.Email
			err := InsertMember(&newMember)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	accessToken, err := jwt.AccessTokenProvider(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
