$(function() {
	function hotSwap(event) {
		var target = event.currentTarget;
		$(".card-container").find("div").each(function(i, elem) {
			$(elem).hide(1000);
		});
		$(target)
			.addClass("expand")
			.show(1000);
	}
	$(".card-container>div").on("click", event, hotSwap);
});