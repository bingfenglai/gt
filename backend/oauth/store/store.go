package store

var ClientStore *ClientDbStore

func init(){

	ClientStore = &ClientDbStore{}

}