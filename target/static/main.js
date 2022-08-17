window.onload = () =>{
  load();
};

function load() {
  let main = document.querySelector("main");
  main.querySelectorAll("a").forEach(a => {
    let oldUrl = a.href;
    a.onclick = e => {
      e.preventDefault();
      naviagate(oldUrl);
    }
    a.href = "#";
  });
}

function naviagate(rawUrl) {
  let oldUrl = new URL(rawUrl);
  let url = new URL(rawUrl);
  url.pathname = "/body"+url.pathname;
  console.log(`Redirecting ${rawUrl} to ${url}`);
  fetch(url).then(r => {
    r.text().then(txt => {
      document.querySelector("main").innerHTML = txt;
      history.pushState(null, null, oldUrl);
      load();
    });
  });
}