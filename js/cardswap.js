$(function() {
	function hotSwap(event) {

		var target = event.currentTarget;
		var parent = $(target).parent("div");

		if ($(parent).hasClass("expand")) {
			target = $(parent);
		}

		$(".card-container").find("div").each(function(i, elem) {
			$(elem).toggle(1000, "linear");
		});

		$(target)
			.toggleClass("expand")
			.show(1000)
			.off("click");

		$("#close-card").toggle();

		if ($(target).hasClass("expand") === false) {
			$(target).on("click", event, hotSwap);
		}
	}

	function closeCard(event) {
		hotSwap(event);
	}

	$(".card-container>div").on("click", event, hotSwap);
	$("#close-card").on("click", event, hotSwap);
	$("#close-card").toggle();
});
