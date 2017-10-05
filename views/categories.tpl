<div>
    <div class="row" style="margin-top: -40px">
        <h2>{{.pageTitle}}</h2>
        <h5>{{.postTotal}} wallpapers found</h5>
        <hr/>
    </div>
    <div class="row">
        {{range $val := .posts}}
        <a href="/wallpaper?id={{$val.Id}}">
            <div class="col-md-4 col-sm-3 col-xs-12 text-center wallpaperTile" style="background: url({{ $val.Thumb}})">
                <p class="wallpaperTileTitle">{{ $val.Title}}</p>
            </div>
        </a>
        {{end}}
    </div>
    <div class="row text-center">
        {{if gt .paginator.PageNums 1}}
        <ul class="pagination pagination-sm">
            {{if .paginator.HasPrev}}
            <li><a href="{{.paginator.PageLinkFirst}}">Prev</a></li>
            <li><a href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
            {{else}}
            <li class="disabled"><a>First</a></li>
            <li class="disabled"><a>&lt;</a></li>
            {{end}}
            {{range $index, $page := .paginator.Pages}}
            <li
                    {{if $.paginator.IsActive .}} class="active" {{end}}>
                <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
            </li>
            {{end}}
            {{if .paginator.HasNext}}
            <li><a href="{{.paginator.PageLinkNext}}">&gt;</a></li>
            <li><a href="{{.paginator.PageLinkLast}}">Next</a></li>
            {{else}}
            <li class="disabled"><a>&gt;</a></li>
            <li class="disabled"><a>Last</a></li>
            {{end}}
        </ul>
        {{end}}
    </div>
</div>

