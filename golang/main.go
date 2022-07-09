package main

import (
  "fmt"
  "io" // Parses HTTP response body
  "log"
  "net/http" // Makes HTTP requests
  "os"
  // "reflect" // Type checking via reflect.TypeOf()
  "strings"
  "text/tabwriter"
  "github.com/droundy/goopt" // Pretty argument parsing
  "github.com/tidwall/gjson" // JSON Parser
)

// ToDo:
// - Handle pagination in API ('.next') and optionally return ALL tags of an image instead of just the latest
// - Return only tags which match the current platform by default, optionally return specific/all OSs and architectures

func httpGet(url, name string) []byte {
  response, err := http.Get(url)
  if err != nil {
    fmt.Println("Failed to call DockerHub API")
    log.Fatal(err)
  }
  responseData, err := io.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Failed to read http response body")
    log.Fatal(err)
  }
  return responseData
}

func main() {

  // Options/Flag Prep
  goopt.Summary = "tag-list DOCKERHUB_REPO"
  goopt.Description = func() string {
    return "Lists available tags for a particular Docker image on Dockerhub. DOCKERHUB_REPO may by a library image repo (eg. 'golang') or a user-created repo (eg. 'electricwarr/tag-list')."
  }
  goopt.Version = "2.0.0"
  goopt.Parse(nil)

  // Validate that there is exactly one positional argument
  if len(goopt.Args) != 1 {
    log.Fatal("You must specify a docker image to list tags of, eg 'golang'. Run 'tag-list --help' for more detail.")
  }
  repo_name := goopt.Args[0]

  // Construct DockerHub URL
  base_url := "https://registry.hub.docker.com/v2/repositories"
  var full_url string
  if strings.Contains(repo_name, "/") {
    full_url = base_url + "/" + repo_name + "/tag"
  } else {
    full_url = base_url + "/" + "library/" + repo_name
  }
  full_url += "/tags"

  // Prepare to store tag info
  type tag_data struct {
    name string
    platforms []string
  }
  tags := make([]tag_data, 0)

  // Call DockerHub API
  response := httpGet(full_url, "DockerHub API")

  // Begin parsing JSON
	results := gjson.Get(string(response), "results")
  // Iterate over query results
  results.ForEach( func(key, tag gjson.Result) bool {
    metadata := new(tag_data)
    tags = append(tags, *metadata)
    thisTag := &tags[len(tags)-1] // uhhh... &tags? Guessed.
    thisTag.name = tag.Get("name").String()
    tag.Get("images").ForEach( func(key, image gjson.Result) bool {
      image_info := image.Map()
      platform := fmt.Sprintf("%s/%s", image_info["os"], image_info["architecture"])
      if image_info["variant"].String() != "" {
        platform = fmt.Sprintf("%s(%s)", platform, image_info["variant"])
      }
      thisTag.platforms = append(thisTag.platforms, platform)
      return true // keep iterating
    } )
    return true // keep iterating
  } )

  // Output Pretty Table
  // Observe how the b's and the d's, despite appearing in the
  // second cell of each line, belong to different columns.
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
  fmt.Fprintln(w, "TAG\tPLATFORMS") // Header
  for _, output_tag := range tags {
    fmt.Fprintln(w, fmt.Sprintf("%s\t%s", output_tag.name, strings.Join(output_tag.platforms, ", ")))
  }
  w.Flush()

}
