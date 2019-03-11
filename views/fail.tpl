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

        .success {
            color: #fc0f1b;
            font-size: 37px;
        }

        .retry {
            background-color: #059056!important;
        }

        .back {
            background-color: #09a5b5!important;
            margin-top: 10%
        }
    </style>
</head>

<body>
    <div class="container">
    <div style="max-width: 480px;
            margin: 27px 33px;"><img src="/static/img/KKP.png" style="width: 100%"></div>
        <div class="content" style="margin-top: 34%">
            <p class="success">فشل إلغاء الاشتراك</p>
            <p class="last">فشل إلغاء الاشتراك ، يرجى المحاولة مرة أخرى</p>
            <a id="retry" type="button" class="unsub retry" onClick="location.href='/unsub/{{.service_id}}'">إعادة المحاولة</a>
            <a id="go_home" type="button" class="unsub back" onClick="location.href='{{.contentURL}}'">اذهب إلى الصفحة الرئيسية </a>
        </div>
    </div>
</body>

</html>
