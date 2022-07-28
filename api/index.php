<?php
  // Start up MySQL.
  include('config/conn.php');
  session_start();

  // Set up the Title.
  echo "<title>Rain API</title>";

  // Function to sanatize all parameters.
  function htmlsc($string) {
    return htmlspecialchars($string, ENT_QUOTES, "UTF-8");
  }

  // Variable to store all keys.
  $keyList = array("");

  // Parameter variables.
  $host = htmlsc($_GET['host']);
  $port = htmlsc($_GET['port']);
  $time = htmlsc($_GET['time']);
  $type = htmlsc($_GET['type']);
  $key = htmlsc($_GET['key']);

  // Error handler.
  if (!isset($_GET["key"]) || !isset($_GET["host"]) || !isset($_GET["port"]) || !isset($_GET["type"]) || !isset($_GET["time"])) {
    $errParameter = fopen("error/parameters.json", "r");
    $read = fread($errParameter, filesize("error/parameters.json"));
    fclose($errParameter);
    die($read);
  }

  $sql = "SELECT * FROM users WHERE APIkey ='$key';";
  $result = mysqli_query($conn, $sql);
  if (mysqli_num_rows($result) === 1) {
    $row = mysqli_fetch_assoc($result);
    if ($row['key'] == $key && $row['Banned'] == false) {
      session_start();
      $_SESSION["loggedin"] = true;
      $_SESSION["id"] = $id;
      $_SESSION["key"] = $userKey;
    }

    else if ($row['key'] == $key && $row['banned'] == true) {
      $errBanned = fopen("error/banned.json", "r");
      $read = fread($errBanned, filesize("error/banned.json"));
      fclose($errBanned);
      die(read);
    }
  }

  else {
    $errInvalid = fopen("error/invalid.json", "r");
    $read = fread($errInvalid, filesize("error/invalid.json"));
    fclose($errInvalid);
    die($read);
  }

  $sql = "SELECT * FROM users WHERE time =$time;";
  $result = mysqli_query($conn, $sql);
  if (mysqli_num_rows($result) === 1) {
    $row = mysqli_fetch_assoc($result);
    if ($time < $row['duration']) {
      $errTime = fopen("error/time.json", "r");
      $read = fread($errTime, filesize("error/time.json"));
      fclose($errTime);
      die($read);
    }
  }

  // API Keys
  $exampleKey = "lol";

  // Here is the part where the methods are actually created.
  if ($type == "EXAMPLE" || $type == "example") {
    $layer7 = "false";
    $exampleMethod = "real";
  }

  else if ($type == "exampleL7" || $type == "exmapleL7") {
    $layer7 = "true";
    $exampleMethod = "real";
  }

  else {
      $err404 = fopen("error/404.json", "r");
      $read = fread($err404, filesize("error/404.json"));
      fclose($err404);
      die($read);
  }
  // Here is the part where the links actually get set so the code could cURL them.
  if ($layer7 == "false") {
    $example = "http://bro.example?key=$key";

    $examplsSend = curl_init();
    curl_setopt($exampleSend, CURLOPT_URL, $example);
    curl_setopt($exampleSend, CURLOPT_RETURNTRANSFER, true);
    $exampleHead = curl_exec($exampleSend);
    curl_close($exampleSend);
    if(!$exampleHead) {
      $errFailed = fopen("error/failed.json", "r");
      $read = fread($errFailed, filesize("error/failed.json"));
      fclose($errFailed);
      die($read);
      return FALSE;
    }

  }

  else if ($layer7 == "true") {
    $example = "http://bro.example?key=$key";

    $examplsSend = curl_init();
    curl_setopt($exampleSend, CURLOPT_URL, $example);
    curl_setopt($exampleSend, CURLOPT_RETURNTRANSFER, true);
    $exampleHead = curl_exec($exampleSend);
    curl_close($exampleSend);
    if(!$exampleHead) {
      $errFailed = fopen("error/failed.json", "r");
      $read = fread($errFailed, filesize("error/failed.json"));
      fclose($errFailed);
      die($read);
      return FALSE;
    }
  }

  die("{ \"error\": false, \"reason\": null, \"Attack info\": { \"target\": \"$host\", \"port\": \"$port\", \"time\": \"$time\" } }");

  // Kill the user session after the attack was done.
  session_unset();
  session_destroy();
?>
