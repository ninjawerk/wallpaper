<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8"/>

    <link rel="shortcut icon" href="favicon.ico" type="image/x-icon">
    <link rel="icon" href="favicon.ico" type="image/x-icon">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta http-equiv="x-pjax-version" content="v123">

    <title>Overwatch Wallpapers</title>

    <meta content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0' name='viewport'/>

    <!--     Fonts and icons     -->
    <link rel="stylesheet" type="text/css"
          href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Roboto+Slab:400,700|Material+Icons"/>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/latest/css/font-awesome.min.css"/>

    <!-- CSS Files -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet"/>
    <link href="static/css/material-kit.css?v=1.1.0" rel="stylesheet"/>
    <link rel='stylesheet' href='static/css/nprogress.css'/>
    <link rel="stylesheet" href="static/css/custom.css">
    <script src="static/js/jquery.min.js" type="text/javascript"></script>

</head>

<body class="index-page">
<div class="header-area" data-parallax="true"
     style="background-image: linear-gradient( rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.5) ), url('static/img/ge90.jpg'); background-size: cover; background-position: center;">
    <nav class="navbar  " role="navigation-demo">
        <div class="container">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" aria-expanded="false">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand weight500" href="/" data-pjax="#pjax-container">
                    Overwatch Gallery
                    <div class="ripple-container"></div></a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="navigation-example-2">
                <ul class="nav navbar-nav   text-center">

                    <li>
                        <a href="/recent">
                            Latest Wallpapers
                        </a>
                    </li>
                </ul>
            </div><!-- /.navbar-collapse -->
        </div><!-- /.container-->
    </nav>

</div>
<div class="main  ">
    <div class="section section-basic">
        <div class="container" id="pjax-container">
{{.LayoutContent}}
        </div>
    </div>
</div>
<footer class="footer">
    <div class="container text-center">
        <div class="copyright ">
            &copy;
            <script>document.write(new Date().getFullYear())</script>
            , made by CocoBeast.
        </div>
    </div>
</footer>
</body>
<!--   Core JS Files   -->
<script src="static/js/bootstrap.min.js" type="text/javascript"></script>
<script src="static/js/material.min.js"></script>

<!--  Plugin for Date Time Picker and Full Calendar Plugin-->
<script src="static/js/moment.min.js"></script>

<!--	Plugin for the Sliders, full documentation here: http://refreshless.com/nouislider/ -->
<script src="static/js/nouislider.min.js" type="text/javascript"></script>

<!--	Plugin for the Datepicker, full documentation here: https://github.com/Eonasdan/bootstrap-datetimepicker -->
<script src="static/js/bootstrap-datetimepicker.js" type="text/javascript"></script>

<!--	Plugin for Select, full documentation here: http://silviomoreto.github.io/bootstrap-select -->
<script src="static/js/bootstrap-selectpicker.js" type="text/javascript"></script>

<!--	Plugin for Tags, full documentation here: http://xoxco.com/projects/code/tagsinput/  -->
<script src="static/js/bootstrap-tagsinput.js"></script>

<!--	Plugin for Fileupload, full documentation here: http://www.jasny.net/bootstrap/javascript/#fileinput -->
<script src="static/js/jasny-bootstrap.min.js"></script>

<!-- Plugin For Google Maps -->

<script src="static/js/atv-img-animation.js" type="text/javascript"></script>

<!-- Control Center for Material Kit: activating the ripples, parallax effects, scripts from the example pages etc -->
<script src="static/js/material-kit.js?v=1.1.0" type="text/javascript"></script>


<!--page loading-->
<script src='static/js/nprogress.js'></script>
<script src='static/js/pjax-standalone.min.js'></script>
<script src="static/js/pageloading.js"></script>

<script type="text/javascript">
    $('.navbar-toggle').click(function () {
        setTimeout(function () {
            $($('.dropdown')[0]).addClass('open')
        }, 500);
    });
    if (location.hash.slice(1) != '') {
        $('#root').load('./parts/' + location.hash.slice(1) + '.html')
    }
    else {
        $('#root').load('./parts/home.html')
    }
    $(window).on('hashchange', function () {
        $('#root').load('./parts/' + location.hash.slice(1) + '.html')
    })
</script>
</html>
