function closeWin() {
    if (navigator.userAgent.indexOf("Firefox") != -1 || navigator.userAgent.indexOf("Chrome") != -1) {
        window.location.href = "about:blank";
        window.close();
    } else {
        window.opener = null;
        window.open("", "_self");
        window.close();
    }
}

function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]); return null;
}


$.ajax({
    type: "GET",
    url: "http://cpx3.allcpx.com/returnid",
    data: {
        type: "video_w",
        affName: GetQueryString("affName"),
        clickId: GetQueryString("clickId"),
        pubId: GetQueryString("pubId"),
        proId: GetQueryString("proId"),
    }
}).done(function (result) {
    ptxid = result
});

$('#submit_btn_mobile').click(function () {
    var clickid = encodeURIComponent("http://cpx3.allcpx.com/subs/getcust/"+ptxid)
    window.location.href = "http://sso.orange.com/mondiamedia_subscription/?method=getcustomer&merchantId=93&langCode=pl&redirect="+clickid
});


