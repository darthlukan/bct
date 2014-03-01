/* main.js
 *
 * @author Brian Tomlinson
 * @contact brian.tomlinson@linux.com
 *
 */

$(function () {
    function defaultContent() {
        $.get("../html/home.html", function(data) {
            $(".container").html(data);
        });
    }

    function notify(event) {
        var section = event.target.id + ".html";

        $.get("../html/" + section, function(data) {
            $(".container").hide(1000, "linear", function() {
                $(".container")
                    .fadeOut(1000)
                    .fadeIn(2000)
                    .html(data);
            });
        });
    }

    defaultContent();

    $("#home").on("click", event, notify);
    $("#projects").on("click", event, notify);
    $("#social").on("click", event, notify);
    $("#random").on("click", event, notify);

    $("#one").fancybox({
        openEffect: 'elastic',
        closeEffect: 'elastic',
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#two").fancybox({
        openEffect: 'elastic',
        closeEffect: 'elastic',
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#three").fancybox({
        openEffect: 'elastic',
        closeEffect: 'elastic',
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#four").fancybox({
        openEffect: 'elastic',
        closeEffect: 'elastic',
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#big").fancybox({
        openEffect: 'elastic',
        closeEffect: 'elastic',
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });
});
