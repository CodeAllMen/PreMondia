
var QualityDatas = {};
var viewDatas = {};
var pageDatas = {};
var subeveryDayData = {};
var data_aff = {};
var subMoData = {};
var  every_date = {
	"columnData":{
		"date":"日期",
		"转化":"",
		"postback回传数":"",
		"花费":"",
		"MT成功数":"",
		"退订":"",
		"留存":"",
		"扣费成功数":"",
		"扣费成功率":"",
		"扣费金额":"美金",
		"Orange收入":"2.773/4.5",
		"Three 收入":"2.622/4.5",
		"T-mobile收入":"2.776/4.5",
		"Ee收入":"2.834/4.5",
		"Vodafone收入":"2.879/4.5",
		"Virgin收入":"1.943/4.5",
		"O2收入":"1.3516/4.5",
		"当天收入":"分成比例",
		"当天":["转化","扣费成功数","退订","花费","收入"],
		"累计":["转化","留存","扣费成功数","扣费成功率","花费","收入","盈亏"],
		"上月累计收入":"XXxx",
		"上月累计花费":"xxxxx",
		"上月盈亏":"xxxxxx",
	}
}


function getAffiliateData() {
	$.ajax({
		url: 'http://cpx3.allcpx.com/aff_data',
		type: 'GET',
		data: data_aff,
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			$.each(data, function (i, c) {
				var pp = [];
				for (var t = 0; t < c.Aff_data.length; t++) {
					for (var a = 0; a < c.Aff_data[t].Ser_list.length; a++) {
						pp.push(c.Aff_data[t].Ser_list[a].Servername);
					}
				}
				console.log(pp);
				aff_html.push('<tr><td rowspan="' + pp.length + '">' + c.AffName + '</td>');
				$.each(c.Aff_data, function (n, q) {
					console.log(q.Ser_list.length + "+b-length");
					aff_html.push('<td rowspan="' + q.Ser_list.length + '">' + q.Pubname + '</td>');
					$.each(q.Ser_list, function (o, p) {
						var activeuser = p.Total_num - p.Unsub_num;
						aff_html.push('<td>' + p.Servername + '</td><td>' + p.Click_num + '</td><td>' + p.Total_num + '</td><td>' + p.PostNum + '</td><td>' + activeuser + '</td><td>' + p.Unsub_num + '</td><td>' + p.SuccessMT_Num + '</td><td>' + p.FailtMT_Num + '</td><td>' + p.Churn_rate + '</td></tr>');
					});
				});
			});
			$("#aff_content").html(aff_html.join(""));
			$("#daochu2").html("<a id='down2' class='btn btn-search' onclick=\"downloadTable2Excal(data_aff.aff_name,data_aff.serverType,data_aff.operator,data_aff.start_time,data_aff.end_time)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});

}

function getMoDateData() {
	$.ajax({
		url: 'http://cpx3.allcpx.com/sub/mo_data',
		type: 'GET',
		data: subMoData,
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			$.each(data, function (i, c) {
				var pp = [];
				for (var t = 0; t < c.ServiceList.length; t++) {
					for (var a = 0; a < c.ServiceList[t].AffSubData.length; a++) {
						pp.push(c.ServiceList[t].AffSubData[a].AffName);
					}
				}
				console.log(pp);
				aff_html.push('<tr><td rowspan="' + pp.length + '">' + c.Operator + '</td>');
				$.each(c.ServiceList, function (n, q) {
					console.log(q.AffSubData.length + "+b-length");
					aff_html.push('<td rowspan="' + q.AffSubData.length + '">' + q.ServiceName + '</td><td rowspan="' + q.AffSubData.length + '">' + q.Price+ '</td>');
					$.each(q.AffSubData, function (o, p) {
						// var activeuser = p.Total_num - p.Unsub_num;
						aff_html.push('<td>' + p.AffName + '</td><td>' + p.SubNum+ '</td><td>' + p.UnsubNum + '</td><td>' + p.PostbackNum + '</td><td>' + p.FailedMt + '</td><td>' + p.SuccessMt + '</td><td>' + p.Amount + '</td></tr>');
					});
				});
			});
			$("#sub_aff_data").html(aff_html.join(""));

			$("#daochu4").html("<a id='down4' class='btn btn-search' onclick=\"downloadTable4Excal(subMoData.aff_name,subMoData.serverType,subMoData.operator,subMoData.start_sub,subMoData.end_sub,subMoData.start_date,subMoData.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});

}


function getQualitypage() {
	$.ajax({
		url: 'http://cpx3.allcpx.com/quality',
		type: 'GET',
		data: QualityDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			console.log(data);
			var aff_html = [];
			$.each(data.data, function (i, c) {
				aff_html.push('<tr><td>' + c.Date + '</td><td>' + c.TotalSubNum + '</td><td>' + c.PostbackNum + '</td><td>' + c.UnsubNum + '</td><td>' + c.ActivateNum + '</td><td>' + c.TotalMt + '</td><td>' + c.RenewNum + '</td><td>' + c.MtFailed + '</td></tr>');
			});
			$("#qm_content").html(aff_html.join(""));
			$("#daochu3").html("<a id='down3' class='btn btn-search' onclick=\"downloadTable3Excal(QualityDatas.aff_name,QualityDatas.pub_id,QualityDatas.serverType,QualityDatas.operator,QualityDatas.sub_date,QualityDatas.end_date)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
};


function GetEveryDaySubPage() {
	var list_title1 =[];
	var affNameColumn =[];
	var clickType =[];
	$.ajax({
		url: 'http://cpx3.allcpx.com/sub/everyday/data',
		type: 'GET',
		data: subeveryDayData,
		dataType: "json",
		success: function (result) {
			var affClickData  = result
			for (var key in every_date.columnData){
				if (key === "转化" || key === "花费" || key === "退订" || key === "postback回传数" || key === "MT成功数"){
					var lens = 1;
					$.each(affClickData.affClickType,function (i,c) {
						lens += c[Object.keys(c)].length
						affNameColumn.push('<th colspan="'+c[Object.keys(c)].length+'">' + Object.keys(c)+ '</th>')
						$.each(c[Object.keys(c)],function (i,key1) {
							clickType.push('<th>'+key1+'</th>')
                		})
       		 		})
					affNameColumn.push('<th>合计</th>')
					clickType.push('<th></th>')
					list_title1.push('<th colspan='+lens+'>'+ key + '</th>');
			}else if (key === "当天" || key === "累计"){
					list_title1.push('<th colspan='+every_date.columnData[key].length+'>'+ key + '</th>');
					$.each(every_date.columnData[key],function (i,c) {
						affNameColumn.push('<th>' + c+ '</th>')
       		 		})
				} else{
					if (key === "上月累计收入"){
						affNameColumn.push('<th style="background-color:#00FF99">'+ affClickData.lastMonthRevenue + '</th>')
					} else if (key === "上月累计花费"){
						affNameColumn.push('<th style="background-color:#00FF99">'+ affClickData.lastMouthSpend +'</th>')
					}else if (key === "上月盈亏"){
						var ProfitAndLoss = ( parseFloat(affClickData.lastMonthRevenue) - parseFloat(affClickData.lastMouthSpend)).toFixed(3)
						affNameColumn.push('<th style="background-color:#00FF99">'+ ProfitAndLoss +'</th>')
					}else{
						affNameColumn.push('<th>' + every_date.columnData[key]+ '</th>')
					}
					list_title1.push('<th>'+ key + '</th>');
					clickType.push('<th></th>')
			}
		}
		for(var i = 1; i <= 9; i++){
    		clickType.push('<th></th>')
		}

	$("#titleName").html('<tr>' + list_title1.join("") +'</tr>' + '<tr>' + affNameColumn.join("") +'</tr>'+'<tr>' + clickType.join("") +'</tr>');

	var lists2 = [];


    $.each(affClickData.data,function (i,c) {
		lists2.push('<tr><td>' + c.Date + '</td>');

		var total_sub = 0     // 转化
		$.each(c.SubData,function (i,subNum) {
			total_sub += parseInt(subNum)
			lists2.push('<td>' + subNum + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + total_sub + '</td>')

		var total_postNum = 0    //  postback回传数
		$.each(c.PostbackData,function (i,postNum) {
			total_postNum += parseInt(postNum)
			lists2.push('<td>' + postNum + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + total_postNum + '</td>')

		var total_postSpend = 0     //  花费
		$.each(c.PostbackSpend,function (i,postSpend) {
			total_postNum += parseInt(postSpend)
			lists2.push('<td>' + postSpend + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + total_postSpend + '</td>')

		var total_Mt_Num = 0   //  MT成功数
		$.each(c.MtNumData,function (i,mt_num) {
			total_Mt_Num += parseInt(mt_num)
			lists2.push('<td>' + mt_num + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + total_Mt_Num + '</td>')

		var total_Unsub = 0    //  退订
		$.each(c.UnSubData,function (i,unsubNum) {
			total_Unsub += parseInt(unsubNum)
			lists2.push('<td >' + unsubNum + '</td>')
        });
		lists2.push('<td style="background-color:#bedead">' + total_Unsub + '</td>')
		lists2.push('<td>' + c.Active + '</td>')
		lists2.push('<td>' + c.SuccessMt + '</td>')
		lists2.push('<td>' + c.MtRate + '</td>')
		lists2.push('<td>' + c.Amout + '</td>')
		lists2.push('<td>' + c.Orange + '</td>')
		lists2.push('<td>' + c.Three + '</td>')
		lists2.push('<td>' + c.Tmobile + '</td>')
		lists2.push('<td>' + c.Ee + '</td>')
		lists2.push('<td>' + c.Vodafone + '</td>')
		lists2.push('<td>' + c.Virgin + '</td>')
		lists2.push('<td>' + c.O2 + '</td>')
		lists2.push('<td>' + c.DayRevenue + '</td>')
		lists2.push('<td style="background-color:#2BD5D5">' + total_sub + '</td>')
		lists2.push('<td style="background-color:#2BD5D5">' + c.SuccessMt + '</td>')
		lists2.push('<td style="background-color:#2BD5D5">' + total_Unsub + '</td>')
		lists2.push('<td style="background-color:#2BD5D5">' + c.DaySpend + '</td>')
		lists2.push('<td style="background-color:#2BD5D5">' + c.DayRevenue + '</td>')

		lists2.push('<td style="background-color:#5EA287">' + c.GrandTotalSub + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.Active + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.TotalSuccessMt + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.GrandTotalMtRate + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.GrandTotalSpend + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.GrandTotalRevenue + '</td>')
		lists2.push('<td style="background-color:#5EA287">' + c.GrandTotalProfitAndLoss + '</td></tr>')
    })
	$("#subtotal").html(lists2.join(""));
	$("#daochu5").html("<a id='down5' class='btn btn-search' onclick=\"downloadTable5Excal(subeveryDayData.date)\">导出表格</a>")

		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
}


$(document).ready(function () {
	$(function () {
		$(".datepicker").datetimepicker({
			showSecond: true,
			showMillisec: true,
			dateFormat: "yy-mm-dd"
		});
	});
	$("#Subscriber").show();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();
});
$("#searchThird").click(function () {
	var start_time = document.getElementById("start_time_aff").value ? document.getElementById("start_time_aff").value : GetSevenDayDate();
	var end_time = document.getElementById("end_time_aff").value ? document.getElementById("end_time_aff").value : NowDate();

	data_aff = {
		"start_time": start_time,
		"end_time": end_time,
		"operator": $("#Opearator option:selected").text(),
		"aff_name": $("#Aff_Name option:selected").text(),
		"serverType": $("#service_type option:selected").text(),
		"clickType": $("#click_type option:selected").text()
	};
	getAffiliateData();
});



$("#searchMoAffDate").click(function () {
	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();

	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date":start_date,
		"end_date":end_date,
		"operator": $("#Opearator2 option:selected").text(),
		"aff_name": $("#Aff_Name2 option:selected").text(),
		"service_type": $("#service_type2 option:selected").text(),
		"clickType": $("#click_type2 option:selected").text(),
	};
	getMoDateData();
});


$("#searchEveryDaySubData").click(function () {
	var monthDate = document.getElementById("sale_month").value ? document.getElementById("sale_month").value : NowMoth();
	subeveryDayData = {
		"date": monthDate
	};
	GetEveryDaySubPage();
});


//change nav page
$("#view_title").click(function () {
	$("#Quality").hide();
	$("#Subscriber").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back1").addClass("background");

});
$("#qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").show();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").hide();

	$("#back1").removeClass("background");
	$("#back4").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back3").addClass("background");

	QualityDatas = {
		"aff_name": "All",
		"operator": "All",
		"date": "2017-09-04"
	};
	getQualitypage();
});
$("#aff_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").show();
	$("#AffSubQuality").hide();
	$("#Everydaysubquality").hide();

	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back5").removeClass("background");
	$("#back2").addClass("background");

	data_aff = {
		"start_time": GetSevenDayDate(),
		"end_time": NowDate(),
		"operator": "All",
		"aff_name": "All",
		"serverType": "All"
	};
	getAffiliateData();
});

$("#sub_qm_title").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#EveryDaySubQuality").hide();
	$("#AffSubQuality").show();


	$("#back3").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").removeClass("background");
	$("#back4").addClass("background");




	var start_sub_date = document.getElementById("start_sub_date").value ? document.getElementById("start_sub_date").value : GetSevenDayDate();
	var end_sub_date = document.getElementById("end_sub_date").value ? document.getElementById("end_sub_date").value : NowDate();
	var start_date = document.getElementById("start_date").value ? document.getElementById("start_date").value : GetSevenDayDate();
	var end_date = document.getElementById("end_date").value ? document.getElementById("end_date").value : NowDate();
	subMoData = {
		"start_sub": start_sub_date,
		"end_sub": end_sub_date,
		"start_date":start_date,
		"end_date":end_date,
		"operator": $("#Opearator2 option:selected").text(),
		"aff_name": $("#Aff_Name2 option:selected").text(),
		"service_type": $("#service_type2 option:selected").text()
	};
	getMoDateData();
});



$("#every_sub_qm").click(function () {
	$("#Subscriber").hide();
	$("#Quality").hide();
	$("#Affiliate").hide();
	$("#AffSubQuality").hide();
	$("#EveryDaySubQuality").show();


	$("#back3").removeClass("background");
	$("#back4").removeClass("background");
	$("#back1").removeClass("background");
	$("#back2").removeClass("background");
	$("#back5").addClass("background");


	GetEveryDaySubPage()
	// getAffiliateData();
});


// view page search
$("#query12").click(function () {
	viewDatas = {
		"operator": $("#telco_view option:selected").text(),
		"start_time": document.getElementById("start_time_view").value ? document.getElementById("start_time_view").value : GetSevenDayDate(),
		"end_time": document.getElementById("end_time_view").value ? document.getElementById("end_time_view").value : NowDate(),
		"aff_name": $("#affName").val(),
		"service_type": $("#serviceType").val(),
		"pubid": $("#pubId option:selected").text(),
		"clickType":$("#clickType option:selected").text()
	};

	$.ajax({
		url: 'http://cpx3.allcpx.com/aff_mt',
		type: 'GET',
		data:viewDatas,
		dataType: "json",
		success: function (result) {
			var data = result;
			var totalMt = data.datas.SuccessMt + data.datas.FailedMt;
			$("#viewtb1").html("<tr><td>" + data.datas.TotalSub + "</td><td>" + data.datas.TotalUnsub + "</td><td>" + data.datas.TotalPostback + "</td><td>" + totalMt + "</td><td>" + data.datas.SuccessMt + "</td><td>" + data.datas.FailedMt + "</td><td>" + data.datas.MtRate + "</td></tr>");
			$("#daochu").html("<a id='down1' class='btn btn-search' onclick=\"downloadTable1Excal(viewDatas.aff_name,viewDatas.pubid,viewDatas.service_type,viewDatas.operator,viewDatas.start_time,viewDatas.end_time)\">导出表格</a>")
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	})
});

// Quality page search


$("#searchQm").click(function () {
	event.preventDefault();
	var time = document.getElementById("Quality_start_time").value;
	time = time.substr(0, 10);
	var end_date = document.getElementById("Quality_end_time").value;
	time = time.substr(0, 10);
	var time = time ? time : GetSevenDayDate();
	var end_date = end_date ? end_date : NowDate();
	QualityDatas = {
		"aff_name": $("#Quality_Aff_Name option:selected").text(),
		"operator": $("#Quality_Opearator option:selected").text(),
		"sub_date": time,
		"end_date": end_date,
		"serverType": $("#Quality_service_type option:selected").text(),
		"pub_id": $("#pud_select option:selected").text(),
		"clickType": $("#Quality_click_type option:selected").text()
	};
	console.log(QualityDatas);
	getQualitypage();
});
// traffic page clear
$("#clearData").click(function () {
	event.preventDefault();
	location.reload();
});

$("#next").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#nav").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#go").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#prev").click(function () {
	event.preventDefault();
	turnTrafficPage();
});
$("#first").click(function () {
	event.preventDefault();
	turnTrafficPage();
});

$("#next1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#nav1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#go1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#prev1").click(function () {
	event.preventDefault();
	turnViewPage();
});
$("#first1").click(function () {
	event.preventDefault();
	turnViewPage();
});



function changeMe1() {
	$.ajax({
		url: 'http://cpx3.allcpx.com/get_pubid?aff_name=' + $("#affName").val(),
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			console.log(aff_html)
			$("#pubId").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
                'selectedText': 'cat'
            });
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}

function changeMe() {
	var a = $("#Quality_Aff_Name").val();
	$.ajax({
		url: 'http://cpx3.allcpx.com/get_pubid?aff_name=' + a,
		type: 'GET',
		dataType: "json",
		success: function (result) {
			var data = result.data;
			console.log(data);
			var aff_html = [];
			aff_html.push('<select class="selectpicker" id="pud_select" data-live-search="true" title="Please select a lunch ..."><option>All</option>')
			$.each(data, function (o, p) {
				aff_html.push('<option>' + p + '</option>');
			});
			aff_html.push('</select>')
			console.log(aff_html)
			$("#ziqudao").html(aff_html.join(""));
			$('.selectpicker').selectpicker({
                'selectedText': 'cat'
            });
		}.bind(this),
		error: function (error) {
			console.log(error);
		}
	});
}



$('.form_datetime2').datetimepicker({
    minView: "month", //选择日期后，不会再跳转去选择时分秒
    language:  'zh-CN',
    format: 'yyyy-mm-dd',
    todayBtn:  1,
    autoclose: 1,
});



function downloadTable1Excal(aff_name,PubId,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"  |PubId:"+PubId+"   |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table1").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down1");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "扣费数量.xls";
}

function downloadTable2Excal(aff_name,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"   |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table2").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down2");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "渠道质量.xls";
}


function downloadTable3Excal(aff_name,PubId,service_type,operator,start_sub, end_sub) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"   |PubId:"+PubId+"      |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"</h4>" + document.getElementById("table3").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文件内容的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down3");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "持续扣费情况.xls";
}


function downloadTable4Excal(aff_name,service_type,operator,start_sub, end_sub,start,end) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>AffName:"+aff_name+"     |ServiceType:"+service_type+"  |Operator:"+operator+"   |Start Sub Date: "+start_sub+" |End Sub Date: "+end_sub+"   |Search Start Date: "+start+"   |Search End Date: "+end+"</h4>" +document.getElementById("table4").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down4");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "订阅留存.xls"
}

function downloadTable5Excal(Month) {
	var html = "<html><head><meta charset='utf-8' /></head><body><h4>"+ Month +"</h4>" +document.getElementById("table5").outerHTML + "</body></html>";
	// 实例化一个Blob对象，其构造函数的第一个参数是包含文的数组，第二个参数是包含文件类型属性的对象
	var blob = new Blob([html], { type: "application/vnd.ms-excel" });
	var a = document.getElementById("down5");
	// 利用URL.createObjectURL()方法为a元素生成blob URL
	a.href = URL.createObjectURL(blob);
	// 设置文件名，目前只有Chrome和FireFox支持此属性
	a.download = "订阅留存.xls"
}


$("#sale_month").datepicker({
 		changeMonth: true,
		changeYear: true,
		dateFormat: 'mm-yy',
		showButtonPanel: true,
		monthNamesShort: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12'],
		closeText: '选择',
		currentText: '本月',
		isSelMon:'true',
		onClose: function (dateText, inst) {
			var month = +$("#ui-datepicker-div .ui-datepicker-month :selected").val() + 1,
				year = $("#ui-datepicker-div .ui-datepicker-year :selected").val();
			if (month < 10) {
				month = '0' + month;
			}
			this.value = year + '-' + month;
			if (typeof this.blur === 'function') {
				this.blur();
			}
		}
});



function GetSevenDayDate() {   //  获取7天前日期
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	myDate.setDate(myDate.getDate() - 7);
	dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
	return dateTemp
}


function NowDate() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1) + "-" + myDate.getDate();
	return dateTemp
}

function NowMoth() {
	var myDate = new Date(); //获取今天日期
	var dateTemp;
	dateTemp = myDate.getFullYear() + "-" + (myDate.getMonth() + 1);
	return dateTemp
}