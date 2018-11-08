package dao

import(
  
)

//get delivery info
func getDelivery(int orderNum, string addr){
  deliv, err := sql.Open("mysql", "")
  if err != nil(
    utils.Logger.Error("Cannot access.")
    return err
  )
  order, err := deliv.Query("select orderNumber, address from delivery where orderNumber = orderNum and address = addr", 1)
  if err != nil(
    return err
  )
  return order
}
