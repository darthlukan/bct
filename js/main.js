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

    function maintenance(event) {
        event.preventDefault();
		alert("This section is under maintenance, check back later!");
    }

    $("#home").on("click", event, render);
    $("#projects").on("click", event, maintenance);
    $("#social").on("click", event, maintenance);
    $("#random").on("click", event, maintenance);
});
