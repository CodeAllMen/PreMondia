package update_mo

// func UpdateMo() {
// 	o := orm.NewOrm()
// 	var notifications []models.Notification
// 	o.QueryTable("notification").OrderBy("id").All(&notifications)
// 	fmt.Println(notifications)
// 	for _, one := range notifications {
// 		var mo = new(models.Mo)
// 		mo.Price = one.Price
// 		mo.Operator = one.Operator
// 		mo.SubscriptionID = one.SubscriptionID
// 		mo.ServiceID = one.ServiceID
// 		mo.CustomerID = one.CustomerID
// 		mo.Channel = one.Channel
// 		mo.PackageCode = one.PackageCode
// 		mo.ProductCode = one.ProductCode
// 		notification.UpdateOrInsertMoTest(one.Action, one.SubscriptionStatus, one.Price, one.Sendtime, one.ID, mo)
// 	}
// }
