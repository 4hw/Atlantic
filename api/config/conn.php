<?php

$sname= "127.0.0.1";
$unmae= "root";
$password = "root";
$db_name = "blissful";
$mertwashere = True;

$conn = mysqli_connect($sname, $unmae, $password, $db_name);

if (!$conn) {
	die("Connection failed!");
}
