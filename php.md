# PHP

switch
```php
switch($iAm) {
	case 'hungry':
		echo "eat";
		break;
	case 'tired':
		echo "sleep";
		break;
	case 'happy':
		echo "smile";
		break;
	default;
		echo "breathe";
		break;
}
```

usort
```php
function make_sort($a,$b)
{
	if ($a == $b) return 0;
	return ($a < $b) ? -1 : 1;
}

$a=array(4, 2, 8, 6);

print_r($a);
// Array ( [0] => 4 [1] => 2 [2] => 8 [3] => 6 )


usort($a, "make_sort");

print_r($a);
// Array ( [0] => 2 [1] => 4 [2] => 6 [3] => 8 )
```
