package  mate

import (
	"fmt"
	"hera"
)

type UserREST struct {
}

//curl 'localhost:8083/Hello/Get?fd=123'
func (this *UserREST) Login(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	
	fmt.Println("[info] sadd  userid , value: "+ phone_number )

	//todo  add user info

	_ , err :=hera.Redis.DoCmd("sadd", "userid", phone_number)
	if err != nil {
		fmt.Println("[warn] sadd  userid error, value: "+ phone_number )
	}

	return c.Success("phone_number : " + phone_number)
}

//curl 'localhost:8083/Hello/Set?fd=123'
func (this *UserREST) MateList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	ret := MatchAlgorithm(phone_number)

	return c.Success(ret)
}

func (this *UserREST) MatedList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	
	ret , err :=hera.Redis.DoCmd("smembers", "like_" + phone_number)
	if err != nil {
		fmt.Println("[warn] sadd  userid error, value: "+ phone_number )
	}

	value, _ := ret.([]string)
//	arr :=  []int{1,2,3}
	//todo get userinfo
	fmt.Printf("%s", ret)
	return c.Success(value)
}

//curl 'localhost:8083/Hello/Set?fd=123'
func (this *UserREST) Like(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	like_id :=  params["like_id"]
	
	fmt.Println("[info] sadd like_"+phone_number+" like_id: "+like_id)

	_ , err :=hera.Redis.DoCmd("sadd", "like_" + phone_number, like_id)
	if err != nil {
		fmt.Println("[warn] sadd  userid error, value: "+ phone_number )
	}
	return c.Success("phone_number : " + phone_number)

}

func (this *UserREST) Unlike(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	like_id :=  params["like_id"]
	
	_ , err :=hera.Redis.DoCmd("sadd", "unlike_" + phone_number, like_id)
	if err != nil {
		fmt.Println("[warn] sadd  userid error, value: "+ phone_number )
	}
	return c.Success("phone_number : " + phone_number)
}

func init() {
	hera.NewRouter().AddRouter(&UserREST{})
}


//login

//getmatelist

//nosure

//sure



