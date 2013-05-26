var morePics=function(){
	$.get("./getPicJson?page="+$("#page").val(),function(data){
		var tmp='<img src="{0}" alt="{1}" width="285px">';
		$.each(data,function(i,pic){
			$("#pics").html($("#pics").html()+tmp.replace('{0}',pic.Url).replace('{1}',pic.Title));
		})
		$("#page").val(parseInt($("#page").val())+1)
	},'json')
	
}