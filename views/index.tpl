<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/css/bootstrap.min.css">
  <style>
    .dropdown-menu {
    max-height: 540px;
    overflow-y: auto;
    }
    img{
      object-fit: cover;
      width: 100%;
      height: 550px;
    }
  </style>
  <title>The Cat Api</title>
  </head>
  <body>
  <div class="container mt-3">
      <div class="row">
        <div class="col-2"></div>
          <div class="col-8">    
            <div class="dropdown dropright">
              <button type="button" class="btn btn-warning dropdown-toggle w-100" data-toggle="dropdown">
                  {{.cat_name}}
              </button>
              <div class="dropdown-menu"> 
                {{range $cat := .cats_category}}
                  <a class="dropdown-item" type="submit" href="/?id={{$cat.ID}}&name={{$cat.Name}}">{{$cat.Name}}</a>
                {{end}}  
              </div>
            </div>
          
            <div class="image mb-3">
              <div id="carouselExampleControls" class="carousel slide" data-ride="carousel">
                <div class="carousel-inner">
                  <div class="carousel-item active">
                    <img class="d-block w-100" src="{{.first_cat.URL}}"  alt="cat images">
                  </div>
                  {{range $cat := .cat_details}}
                    <div class="carousel-item">
                      <img class="d-block w-100" src="{{$cat.URL}}"  alt="cat images">    
                    </div>
                  {{end}}
                </div>
                <a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
                  <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                  <span class="sr-only">Previous</span>
                </a>
                <a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
                  <span class="carousel-control-next-icon" aria-hidden="true"></span>
                  <span class="sr-only">Next</span>
                </a>
              </div>
            </div>
            <div class="text-center">
              <h1>{{.cname}}</h1>
              <p><b>ID: {{.id}}</b></p>
              <p>{{.description}}</p>
              <p>--------</p>
              <p><i>{{.temparent}}</i> <br>
                  {{.origin}} <br>
                  {{.weight}} kgs <br>
                  {{.lifespan}} average lifespan
              </p>
              <a class="text-warning" href="{{.wikipedia}}">WIKIPEDIA</a>
            </div>
          </div>
        <div class="col-2">
        </div>
      </div>
    </div>
  <script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.slim.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/js/bootstrap.bundle.min.js"></script>
  </body>
</html>
