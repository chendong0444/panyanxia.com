var morePics=function(){
	$.get("./getPicJson?page="+$("#page").value,function(data){
		var tmp='<img src="{0}" alt="{1}" width="200px">';
		$.each(data,function(i,pic){
			$("#pics").html($("#pics").html()+tmp.replace('{0}',pic.Url).replace('{1}',pic.Title));
		})
		$("#page").value(ParseInt($("#page").value)+1)
	},'json')
	
}