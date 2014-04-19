//Set the max size of the panels to all the panels


$(document).ready(function(){
	
	var maxHeight = 0
	$(".panel").each(function(index,panel){
		if (panel.clientHeight > maxHeight){
			maxHeight = $(this).height();
		}
	});

	$(".panel").each(function(index,panel){
		$(this).height(maxHeight);
	});
});