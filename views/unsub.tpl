<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=360">
    <title>Unsubscribe</title>
    <link type="text/css" rel="stylesheet" href="/static/unsub/style.css">
    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
    <style type="text/css">
        .ver {
            font-size: 12px;
            color: #1f1f1f;
        }
    </style>
</head>

<body>
    <div class="container">
    <center><div style="    font-size: 33px;
                            margin: 29px 0 -65px 0px;
                            color: #000;">إلغاء الاشتراك</div></center>
        <div class="content">
            <form action="unsubPin" id="form" method="Post">
                <p class="m-b">يرجى إدخال رقم هاتفك: </p>
                <input class="fill_num m-b" name="msisdn" id="msisdn" placeholder="20xxxxxxxxx" />
                 <input class="noexit" id="service_id" name="service_id" value="{{.service_id}}" style="display: none;"/>
                <p class="ver">سيتم إرسال رمز تحقق إلى رقم الهاتف هذا..</p>
                <a type="button" class="unsub" id="unsub">عرض</a>
                <p class="last">*الرجاء إدخال رقمك مع رمز البلد (على سبيل المثال 20 في  EG)</p>
            </form>
        </div>
    </div>
</body>
<script>
    if ({{.error}}=="0"){
        alert("رقم الهاتف غير موجود. يرجى المحاولة مرة أخرى.")
    }
    $("#unsub").click(function () {
        if ($("#msisdn").val() == "" || $("#msisdn").val().length < 9||$("#msisdn").val().length>11) {
            alert("Phone number invalid")
        } else if ($("#msisdn").val().length ==9){
            $("#msisdn").val("48"+$("#msisdn").val())
            $("#form").submit()
        }else{
            $("#form").submit()
        }
    });
</script>

</html>
