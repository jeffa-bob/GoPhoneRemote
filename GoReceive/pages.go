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

<body>
    <div class="mdl-textfield mdl-js-textfield getmdl-select">
        <input type="text" value="" class="mdl-textfield__input" id="devicebox" readonly>
        <input type="hidden" value="" name="devicebox">
        <label for="devicebox" class="mdl-textfield__label">Device</label>
        <ul for="devicebox" class="mdl-menu mdl-menu--bottom-left mdl-js-menu">
            <li class="mdl-menu__item" data-val="DEU">Germany</li>
            <li class="mdl-menu__item" data-val="BLR">Belarus</li>
            <li class="mdl-menu__item" data-val="RUS">Russia</li>
        </ul>
    </div>
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
