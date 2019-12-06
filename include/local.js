var client = new ClientJS();
var fingerprint = client.getFingerprint(); // Calculate Device/Browser Fingerprint
console.log(fingerprint);
document.write('<form action="/report" method="post">');
document.write("  <fieldset>");
document.write(
	'    <input type="text" id="fingerprint" name="fingerprint" readonly="true" value="' +
		fingerprint +
		'">'
);
document.write("     Fingerprint");
document.write("    </input>");
document.write('    <input type="submit" value="Submit">');
document.write("  </fieldset>");
document.write("</form>");
