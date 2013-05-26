var morePics=function(){
	$.get("./getPicJson?page="+$("#page").val(),function(data){
		var tmp='<img src="{0}" alt="{1}" width="285px">';
		$.each(data,function(i,pic){
			$("#pics").append(tmp.replace('{0}',pic.Url).replace('{1}',pic.Title));
		})

		if(data && data.length>1) $("#page").val(parseInt($("#page").val())+1)
	},'json')
	return false;
}