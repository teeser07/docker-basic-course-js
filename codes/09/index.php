<?php 
if(isset($_POST["submit"])) {
    move_uploaded_file($_FILES['file']['tmp_name'], 'files/'.basename($_FILES['file']['name']));
    echo '<h2>Upload done!</h2>';
}
?>
<form action="index.php" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit" value="Upload file" name="submit">
</form>