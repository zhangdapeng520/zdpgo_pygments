<?php
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
// 定义变量并默认设置为空值
$name = $email = $gender = $comment = $website = "";








// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值// 定义变量并默认设置为空值
if ($_SERVER["REQUEST_METHOD"] == "POST")
{
  $name = test_input($_POST["name"]);
  $email = test_input($_POST["email"]);
  $website = test_input($_POST["website"]);
  $comment = test_input($_POST["comment"]);
  $gender = test_input($_POST["gender"]);
}

function test_input($data)
{
  $data = trim($data);
  $data = stripslashes($data);
  $data = htmlspecialchars($data);
  return $data;
}
?>