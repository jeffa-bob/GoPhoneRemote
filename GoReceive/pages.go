package main

const (
	startform = `
<html>

<head>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.green-blue.min.css" />
    <link rel="stylesheet" href="http://fonts.googleapis.com/css?family=Roboto:300,400,500,700" type="text/css">
    <script defer src="https://code.getmdl.io/1.3.0/material.min.js"></script>
</head>

<body><style>
.demo-list-icon {
  width: 300px;
}
</style>

<ul class="demo-list-icon mdl-list">
  <li class="mdl-list__item">
    <span class="mdl-list__item-primary-content">
    <i class="material-icons mdl-list__item-icon">person</i>
    Bryan Cranston
</span>
  </li>
  <li class="mdl-list__item">
    <span class="mdl-list__item-primary-content">
    <i class="material-icons mdl-list__item-icon">person</i>
    Aaron Paul
  </span>
  </li>
  <li class="mdl-list__item">
    <span class="mdl-list__item-primary-content">
    <i class="material-icons mdl-list__item-icon">person</i>
    Bob Odenkirk
  </span>
  </li>
</ul>
    <form action="#">
        <div style="margin-left: 2%" class="mdl-textfield mdl-js-textfield">
            <input class="mdl-textfield__input" type="text"
                pattern="^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"
                id="portinput">
            <label class="mdl-textfield__label" for="portinput">Port</label>
            <span class="mdl-textfield__error">Input is not a Port!</span>
        </div>
    </form>
    <button onclick="submit" style="margin-left: 2%; width: 94%;"
        class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent">
        Submit
    </button>
    <script>
        function submit() {
            var ip = document.getElementById("ipinput");
            var port = document.getElementById("portinput");
            if (!ip.checkValidity() || !port.checkValidity()) {
                return false;
            }

        }
    </script>
</body>

</html>
`
)
