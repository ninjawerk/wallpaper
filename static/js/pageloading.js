/**
 * Created by desha on 10/4/2017.
 */
$(document).pjax('[data-pjax] a, a[data-pjax]', '#pjax-container')
$.pjax.defaults.timeout = 6000;

$('#pjax-container').on('pjax:start', function() {
    NProgress.start();
});

$('#pjax-container').on('pjax:end',   function() {
    NProgress.done();
    $('#pjax-container').fadeOut(0);
    $('#pjax-container').fadeIn(500);
});
