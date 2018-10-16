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
    <center><div style="font-size: 27px;
    margin: 29px 0 -65px 0px;">Anuluj subskrypcję</div></center>
        <div class="content">
            <form action="unsubPin" id="form" method="Post">
                <p class="m-b">Proszę podać swój numer telefonu: </p>
                <input class="fill_num m-b" name="msisdn" id="msisdn" placeholder="48xxxxxxxxx" />
                <p class="ver">Kod weryfikacyjny zostanie wysłany do tego msisdn.</p>
                <a type="button" class="unsub" id="unsub">Prześlij</a>
                <p class="last">*Proszę podać swój numer z kodem kraju (np. 48 w PL)</p>
            </form>
        </div>
    </div>
</body>
<script>
    if ({{.error}}=="0"){
        alert("Nie znaleziono numeru telefonu. Spróbuj ponownie. \n\n Jeśli problem dotyczy presist, skontaktuj się z obsługą klienta pod adresem support@mondiapay.com)
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