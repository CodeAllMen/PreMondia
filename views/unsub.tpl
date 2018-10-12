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
    </style>
</head>

<body>
    <div class="container">
        <div class="content">
            <form action="unsubPin" id="form" method="Post">
                <p class="m-b">Mobile number to be unsubscribed</p>
                <input class="fill_num m-b" name="msisdn" id="msisdn" placeholder="48xxxxxxxxx" />
                <p class="ver">Verification Code will be sent to this msisdn.</p>
                <a type="button" class="unsub" id="unsub">Submit</a>
                <p class="last">*Please enter your number with country code (e.g.48 in the PL)</p>
            </form>
        </div>
    </div>
</body>
<script>
    if ({{.error}}=="0"){
        alert("Phone number not found. Please retry. \r\n If problem presist, kindly contact our customer support at support@mondiapay.com")
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