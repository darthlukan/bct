$(function() {
	function hotSwap(event) {
		var target = event.currentTarget;
		$(".card-container").find("div").each(function(i, elem) {
			$(elem).toggle(1000, "linear");
		});
		$(target)
			.toggleClass("expand")
			.show(1000);
	}
	$(".card-container>div").on("click", event, hotSwap);
});