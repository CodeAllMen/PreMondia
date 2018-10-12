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
        <div class="content">
            <form action="/getCust" id="form" method="Post">
                <p class="m-b">MSISDN:{{.phone}}</p>
                <input class="fill_num m-b" id="pin" name="pin" placeholder="PIN : XXX" />
                <input class="noexit" id="id" name="id" value="{{.id}}"/>
                <p class="ver">Enter PIN receive in SMS</p>
                <a type="button" class="unsub" id="ver">Verify</a>
            </form>
        </div>
    </div>
</body>
<script>
    $("#ver").click(function () {
        if ($("#pin").val() == "") {
            alert("please input crecot pin code!")
        } else {
            $("#form").submit();
        }
    });
    if ({{.error}} == "201"){
        alert("Your Pin Code is wrong.Please check it!")
    }
</script>
</html>