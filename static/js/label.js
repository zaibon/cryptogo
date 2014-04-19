//display the positive transctin in green and
//the negative in orange
$(document).ready(function(){
	$(".label").each(function(i,val){
		var amount = parseFloat(val.innerHTML)
	    if( amount > 0){
	        val.classList.add("label-success");
	    }else if (amount < 0){
	    	val.classList.add("label-warning");
	    }else{
	    	val.classList.add("label-primary");
	    }
	});
});