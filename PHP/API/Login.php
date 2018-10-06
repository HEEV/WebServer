<?php

session_start();

$u = strtolower($_POST['Username']);
$p = strtolower($_POST['Password']);

if ($u == "cedarville" && $p == "bebold") {
  $_SESSION['LoggedIn'] = TRUE;
  header("Location: ../index.php");
} else {
  echo "Invalid login information!";
}
