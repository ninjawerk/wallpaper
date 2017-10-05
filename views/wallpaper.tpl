<div>

    <img src="{{.post.DisplayImage}}" style="width: 100%; height: auto;">
    <h2>{{.post.Title}}</h2>
    <hr/>
    <h4>Download links</h4>
    {{range $size := .post.Sizes}}
    <a download href="{{$size.Url}}" id="{{$size.Size}}" class="btn btn-default btn-sm btn-round">{{$size.Size}}</a>
    {{end}}
    <hr/>
    <h4>Tags</h4>
    <p id="yourRes"></p>
    <div class=" form-group" style="margin-top: -10px">
        <input name="tags" data-role="tagsinput" class="tagsinput addborder tag-rose form-control" value="{{.tags}}"
               disabled/>
    </div>
    <script>
        var h = window.screen.height;
        var w = window.screen.width;
        $("#" + w +"x" +h).addClass("btn-primary ")
        $("#" + w +"x" +h).removeClass("btn-default ")
    </script>
</div>

