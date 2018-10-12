<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>首页</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no"	name="viewport">
    <link type="text/css" rel="stylesheet" href="../static/css/bootstrap.min.css" />
    <link type="text/css" rel="stylesheet" href="../static/css/gameCss.css" />
    <script type="text/javascript" src="../static/js/jquery.min.js" ></script>
    <script type="text/javascript" src="../static/js/bootstrap.min.js"></script>
    <script type="text/javascript"  src="../static/js/indiaIndex.js"></script>
    <script>
        $(function(){
            $(function(){
            alert({{.first}});
                        if({{.first}}=="yes"){
                $("#urlAlert").removeAttr("hidden")
            }else{
                $("#urlAlert").attr("hidden","hidden");
            }
        });

        $("#closeAlert").click(function(){
            $("#urlAlert").attr("hidden","hidden");
        })
        });

    </script>

    <!-- Global Site Tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-106568989-1"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag(){dataLayer.push(arguments)};
        gtag('js', new Date());

        gtag('config', 'UA-106568989-1');
    </script>

</head>
<body>
<!--header logo和登录按钮 -->
<header>
    <div class="container">
        <div class="col-md-3 col-sm-4 col-xs-6 logo">
            <img src="../static/img/gamehub.png" class="logoImg" id="logoImg" />
        </div>
        <div class="sign" id="sign">
            <div class="signTwo">
                <a href="#" class="btn showAll smallScreen" onclick="personal()">
                <i class="icon icon-user"></i>
                 <span id="user_name">{{.username}}</span>
                </a>
                <div class="smallPTwo modal-bodyText line"><i class="icon icon-money"></i>
                    <div class="smallPTwo paddingTop"><span class="moneyFont">×</span>
                        <lable class="moneyFont">{{.money}}</lable></div></div></div>
        </div>
    </div>
</header>
<!-- 页面的导航 -->
<nav class="navColor">
    <div class="container">
        <ul class="navM">
            <li>
                <a href="#" id="home" class="a-font">HOME</a>
            </li>
            <li>
                <a href="#" class="a-font" id="games" >GAMES<span class="caret"></span></a>
            </li>
        </ul>
    </div>
    <div class="nav-dropdown">
        <div class="container">
            <div class="row paddingLeft paddingRight">
                <div class="col-sm-3 col-md-3 noPadding daohang">
                    <ul class="list-group">
                        <li><a href="#" class="a-font" onclick="gameMore('best')" >Best Games</a></li>
                        <li><a href="#" class="a-font" onclick="gameMore('new1')" >New Games</a></li>
                    </ul>
                    <div class="col-sm-7 col-xs-8" >
                        <img src="../static/img/monkey.png" class="monkey" />
                    </div>
                </div>
                <div class="col-sm-9 col-md-9 noPadding">
                    <ul class="list-group appendLi">

                    </ul>
                </div>
            </div>
        </div>
    </div>
</nav>

<div id="body">

</div>

<!--play游戏遮罩开始 -->
<div class="modal fade" id="modal-container-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
            </div>
            <div class="modal-body modal-bodyText">
                <h3 class="h3">Make sure to buy this game ?</h3>
                <div class="modal-bodyDiv">
                    <div class="smallPOne">Price</div>
                    <div class="icon icon-money smallPOne"></div>
                    <div class="smallPOne moneyFont" id="moneyPrice">×12</div>
                    <span class="line" >|</span>
                    <div class="smallPTwo moneyFont" id="userMoneyPrice">×12</div>
                    <div class="icon icon-money smallPTwo"></div>
                    <div class="smallPTwo">coins</div>
                    <input type="hidden" id="game_idDemo" />
                </div>
                <h5>The game you have purchased will be added to Personal Center</h5>
                <h6 class="coinsError" id="coinsError" >

                </h6>
            </div>
            <div class="modal-footer">
                <!--<button type="button" class="btn showAll" id="demoGame">Demo</button><button type="button" class="btn playNow" id="purchase" >Purchase</button>-->
                <button type="button" class="btn playNow" id="purchase" >Yes</button><button type="button" data-dismiss="modal" class="btn showAll">No</button>
            </div>
        </div>
    </div>
</div>
<!--url -->
<div class="urlAlert" id="urlAlert" hidden="hidden">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header borderHeader">
                <h4 class="modal-title" id="myModalLabel">
                    Please Remember This URL:
                </h4>
            </div>
            <div class="modal-body modal-bodyText">

                <h5>
                    {{.url}}
                </h5>
            </div>
            <div class="modal-footer">
                  <button type="button" data-dismiss="modal" class="btn showAll"  id="closeAlert">Yes</button>
            </div>
        </div>
    </div>
</div>
</body>
</html>