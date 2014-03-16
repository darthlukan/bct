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
		event.preventDefault();
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

    function eventCheck(event) {
        console.log(event);
    }

    $("#home").on("click", event, render);
    $("#projects").on("click", event, render);
    $("#social").on("click", event, render);
    $("#random").on("click", event, render);
});
