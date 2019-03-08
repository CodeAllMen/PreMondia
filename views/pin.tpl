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
            color: #a1a1a1;
        }
        .noexit {
            display: none;
        }
    </style>
</head>

<body>
    <div class="container">
     <div style="max-width: 480px;
        margin: 27px 33px;"><img src="/static/img/KKP.png" style="width: 100%"></div>
        <div class="content">
            <form action="/getCust" id="form" method="Post">
                <p class="m-b">MSISDN: {{.phone}}</p>
                <input class="fill_num m-b" id="pin" name="pin" placeholder="PIN : XXX" />
                <input class="noexit" id="id" name="id" value="{{.id}}" style="display: none"/>
                <input class="noexit" id="service_id" name="service_id" value="{{.service_id}}" style="display: none"/>
                <p class="ver">أدخل رمز PIN المستلم في رسالة SMS</p>
                <a type="button" class="unsub" id="ver">تحقق من</a>
            </form>
        </div>
    </div>
</body>
<script>
    $("#ver").click(function () {
        if ($("#pin").val() == "") {
            alert("Proszę podać swój PIN!")
        } else {
            $("#form").submit();
        }
    });
    if ({{.error}} == "201"){
        alert("Twój PIN jest nieprawidłowy. Spójrz!")
    }
</script>
</html>
