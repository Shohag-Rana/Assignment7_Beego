package controllers

import (
	ud_func "catapi/controllers/requiredfunction" //ud_func = user define function
	ud_struct "catapi/controllers/structure" //ud_struct = user define structure
	"encoding/json"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/patrickmn/go-cache"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	cats_category_channel_data := make(chan string)
	go ud_func.Response_Url("https://api.thecatapi.com/v1/breeds", cats_category_channel_data)
	data := <-cats_category_channel_data
	var cats_category = ud_struct.Cats_Category{}
	err := json.Unmarshal([]byte(data), &cats_category)
	if err != nil {
		fmt.Println("Can't unmarshal JSON object into struct")
	}
	c.Data["cats_category"] = cats_category

	// catch the cats category id and name
	var id, name string
	c.Ctx.Input.Bind(&id, "id")
	c.Ctx.Input.Bind(&name, "name")
	if id != "" {
	} else {
		id = "beng" // default cats type bengal
		name = "Bengal"
	}

	// get the cats category details  with cats id
	api_key := ud_func.GoDotEnvVariable("API_KEY")
	cats_details_channel_data := make(chan string)
	go ud_func.Response_Url("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids="+id+api_key, cats_details_channel_data)

	var cat_details = ud_struct.Cats_Category_Details{}
	err = json.Unmarshal([]byte(<-cats_details_channel_data), &cat_details)
	if err != nil {
		fmt.Println("Can't unmarshal JSON object into struct")
	}

	// cache operation
	cs := cache.New(25*time.Minute, 40*time.Minute)
	cs.Set(id, cat_details, cache.DefaultExpiration)

	x, found := cs.Get(id)
	if found {
		fmt.Println(x)
	}

	c.Data["cat_details"] = cat_details[1:] // all the cats details exclude first one for slide show
	c.Data["first_cat"] = cat_details[0]    // extract the first cats in the list for view in the first slide
	c.Data["cat_name"] = name
	//Cats category details data

	c.Data["weight"] = (cat_details[0].Breeds[0].Weight.Imperial)
	c.Data["id"] = (cat_details[0].Breeds[0].ID)
	c.Data["temparent"] = (cat_details[0].Breeds[0].Temperament)
	c.Data["cname"] = (cat_details[0].Breeds[0].Name)
	c.Data["origin"] = (cat_details[0].Breeds[0].Origin)
	c.Data["description"] = (cat_details[0].Breeds[0].Description)
	c.Data["lifespan"] = (cat_details[0].Breeds[0].LifeSpan)
	c.Data["wikipedia"] = (cat_details[0].Breeds[0].WikipediaURL)

	c.TplName = "index.tpl" //template(user interface)
}
