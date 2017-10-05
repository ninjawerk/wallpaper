<div class="row" style="margin-top: -40px">
    <h2 class="" style="text-transform: capitalize">Browse the best overwatch Wallpapers!</h2>
    <h3 class="">Zero Ads, Only Wallpapers</h3>
    <hr/>
    <div  style="  text-align: center; position: relative" >
        {{range $val := .cats}}
        <a href="/category?id={{$val.Title}}"  data-pjax="#pjax-container" >
            <div class="col-md-2 col-sm-3 col-xs-12 text-center categoryTile" style="background-image: url({{ $val.Thumb}})">
                <p class="categoryTileTitle">{{ $val.Title}}</p>
            </div>
        </a>
        {{end}}
    </div>
</div>

