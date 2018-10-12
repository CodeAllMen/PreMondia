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
    <script type="text/javascript"  src="../static/css/game.js"></script>
</head>
<body>
<!--header logo和登录按钮 -->
<header>
    <div class="container">
        <div class="col-md-3 col-sm-4 col-xs-6 logo">
            <img src="../static/img/gamehub.png" class="logoImg" id="logoImg" />
        </div>
        <div class="sign" id="sign">

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
<!--登录按钮的遮罩层 -->
<div class="modal fade" id="modal-container-801879" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
            </div>
            <form id="formOne" action="" >
                <div class="modal-body noPadding">

                    <div class="row">
                        <div class="col-md-6 col-sm-6 col-sm-offset-3 col-xs-8 col-xs-offset-2">
                            <img src="../static/img/gamehub.png" class="logoImg" />
                        </div>
                        <div class="col-md-8 col-sm-8 col-sm-offset-2 col-xs-8 col-xs-offset-2">
                            <div class="zhezhao">
                                <input type="text" class="form-control radius" id="userNameOne"  placeholder="UserName">
                            </div>
                            <div class="zhezhao" >
                                <input type="password" class="form-control radius" id="userPasswordOne"  placeholder="Password">
                            </div>
                            <div class="zhezhao" >
                                <div class="radioT">
                                    <input type="radio" name="deo" id="deo1" >
                                    <label for="deo1" id="deo1L" onMouseOver="this.name=document.getElementById('deo1').checked" ></label>
                                </div>
                                <div style="float:left">
                                    <span class="radioLable">Remember This Account</span>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
                <div class="modal-footer">
                    <!--<button type="button" class="btn showAll" id="userLogin" >LOGIN</button><a id="modal-801870" href="#modal-container-801870" role="button" data-toggle="modal" class="btn gray">SIGN UP</a>-->
                    <a  href="#" role="button" data-toggle="modal" id="checkFirstLogin" class="btn showAll">LOGIN</a><button type="button" class="btn gray" >XX</button>
                </div>
            </form>
        </div>
    </div>
</div>
<div class="modal fade" id="modal-container-801870" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
            </div>
            <div class="modal-body noPadding">
                <div class="row">
                    <div class="col-md-6 col-sm-6 col-sm-offset-3 col-xs-8 col-xs-offset-2">
                        <img src="../static/img/gamehub.png" class="logoImg" />
                    </div>
                    <div class="col-md-8 col-sm-8 col-sm-offset-2 col-xs-8 col-xs-offset-2">
                        <!--<div class="zhezhao">
                            <input type="text" class="form-control radius"  id="userNameTwo"  placeholder="UserName">
                        </div>
                        <div class="zhezhao">
                            <input type="text" class="form-control radius" id="userEmail"  placeholder="Email">
                        </div>
                        <div class="zhezhao" >
                            <input type="text" class="form-control radius" id="userEmailTwo"  placeholder="Confirm Email">
                        </div>
                        <div class="zhezhao" >
                            <input type="password" class="form-control radius" id="userPasswordTwo"  placeholder="Password">
                        </div>
                        <div class="zhezhao" >
                            <div class="radioT">
                                <input type="radio" name="deo1" id="deo2" >
                                <label for="deo2" id="deo2L" onMouseOver="this.name=document.getElementById('deo1').checked" ></label>
                            </div>
                            <div style="float:left">
                                <span class="radioLable">Remember This Account</span>
                            </div>
                        </div>-->
                        <div class="zhezhao">
                            <input type="text" class="form-control radius"  id="passwordOne"  placeholder="Set New Password">
                            <input type="hidden" class="form-control radius"  id="userNameTwo" >
                        </div>
                        <div class="zhezhao">
                            <input type="text" class="form-control radius"  id="passwordTwo"  placeholder="Confirm New Password">
                        </div>
                        <div class="zhezhao">
                            <input type="text" class="form-control radius"  id="emailOne"  placeholder="Set Email">
                        </div>
                        <div class="zhezhao">
                            <h5>We will give you 6 coins for each successful payment every day until you cancel your subscription.</h5>
                        </div>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <!--<button type="button" id="signUp" class="btn playNow" >SIGN UP</button> <button type="button" data-dismiss="modal" class="btn gray">CLOSE</button>-->
                <button type="button" id="setInformation" class="btn playNow" >Yes</button>
            </div>
        </div>

    </div>
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
                    <div class="smallPTwo">Stock</div>
                    <input type="hidden" id="game_idDemo" />
                </div>
            </div>
            <div class="modal-footer">
                <!--<button type="button" class="btn showAll" id="demoGame">Demo</button><button type="button" class="btn playNow" id="purchase" >Purchase</button>-->
      <button type="button" class="btn playNow" id="purchase" >Yes</button><button type="button" data-dismiss="modal" class="btn showAll">No</button>
            </div>
        </div>
    </div>
</div>
<!--play 游戏遮罩结束 -->
<script>

    $("#deo1").click(function(){
        if(document.getElementById('deo1L').name==true){
            this.checked=false;
        }else{
            this.checked=true;
        };
    });
    $("#deo2").click(function(){
        if(document.getElementById('deo2L').name==true){
            this.checked=false;
        }else{
            this.checked=true;
        };
    });
</script>
</body>
</html>