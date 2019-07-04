package html

const NotFoundError =
`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>页面出错了</title>

    <style>
        body {
            background-color: #fff;
        }
        .auto {
            width: 1000px;
            margin: 230px auto;
        }
        .container {
            background: url(/static/images/error-bg.png) no-repeat 560px top;
        }
        .settings {
            padding-left: 200px;
            padding-bottom: 50px;
        }
        .settings .icon {
            display: block;
            width: 242px;
            height: 106px;
            background: url(/static/images/404.png) no-repeat 0 0;
        }
        .settings h4 {
            margin: 30px 0 15px 0;
            font-size: 18px;
            color: #2cb7fd;
        }
        .settings p {
            font-size: 14px;
            color: #999;
        }
        .settings > div {
            margin-top: 40px;
            font-size: 0;
        }
        .settings > div a {
            display: inline-block;
            padding: 10px 40px;
            border: 1px solid #2cb7fd;
            font-size: 15px;
            color: #2cb7fd;
            text-decoration: none;
        }
        .settings > div a:first-child {
            margin-right: 20px;
            color: #fff;
            background-color: #2cb7fd;
        }
    </style>

</head>
<body>
<div class="auto">
    <div class="container">
        <div class="settings">
            <i class="icon"></i>
            <h4>{[.content]}！</h4>
        </div>
    </div>
</div>
</body>
</html>`