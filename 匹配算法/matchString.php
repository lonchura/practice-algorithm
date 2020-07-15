<?php
$a = 'abccdbc';
$b = 'bcc';
$result = match($b, $a);
var_dump($result);

/**
 * @param $pattern
 * @param $str
 * @return array
 */
function match($pattern, $str) {
    $letter = $pattern[0];

    $match = [];
    for($i=0; $i<strlen($str); $i++) {
        if($letter == $str[$i]) {
            if(strlen($pattern) == 1) {
                $match[] = substr($str, 0, $i+1);
            } else {
                $list = match(substr($pattern, 1), substr($str, $i+1));
                foreach($list as $part) {
                    $match[] = sprintf("%s%s", substr($str, 0, $i+1), $part);
                }
            }
        }
    }

    return $match;
}