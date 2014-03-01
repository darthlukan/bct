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

    function render(event) {
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

    $("#home").on("click", event, render);
    $("#projects").on("click", event, render);
    $("#social").on("click", event, render);
    $("#random").on("click", event, render);

    $("#one").fancybox({
        afterClose: function () {
            $("#one").show();
        },
        openEffect: 'elastic',
        closeEffect: 'elastic',
        width: "70%",
        height: "70%",
        maxWidth: "880",
        maxHeight: "600",
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#two").fancybox({
        afterClose: function () {
            $("#two").show();
        },
        openEffect: 'elastic',
        closeEffect: 'elastic',
        width: "70%",
        height: "70%",
        maxWidth: "880",
        maxHeight: "600",
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#three").fancybox({
        afterClose: function () {
            $("#three").show();
        },
        openEffect: 'elastic',
        closeEffect: 'elastic',
        width: "70%",
        height: "70%",
        maxWidth: "880",
        maxHeight: "600",
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#four").fancybox({
        afterClose: function () {
            $("#four").show();
        },
        openEffect: 'elastic',
        closeEffect: 'elastic',
        width: "70%",
        height: "70%",
        maxWidth: "880",
        maxHeight: "600",
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });

    $("#big").fancybox({
        afterClose: function () {
            $("#big").show();
        },
        openEffect: 'elastic',
        closeEffect: 'elastic',
        width: "70%",
        height: "70%",
        maxWidth: "880",
        maxHeight: "600",
        helpers: {
            overlay: {
                css: {
                    'background': 'rgba(0, 0, 0, 0.8)'
                }
            }
        }
    });
});
