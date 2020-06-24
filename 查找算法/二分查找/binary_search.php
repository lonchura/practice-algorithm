<?php
$item = 20;
$list = [9, 10, 20];

$result = binary_search($list, $item);
printf("List: %s\n", json_encode($list));
printf("Search Item: (%s)\n", $item);
printf("Search\n  - index: %s\n  - count: %s\n", $result[0], $result[1]);

function binary_search(array $list, $item) {
    // init pos
    $low = 0;
    $high = count($list) - 1;

    // search
    $_scount = 0;
    while($low <= $high) {
        $_scount++; // increase search count

        $mid = ceil(($low + $high) / 2);
        $guess = $list[$mid];

        if ($guess == $item) {
            return [$mid, $_scount];
        } else if ($guess < $item) {
            $low = $mid + 1;
        } else {
            $high = $mid - 1;
        }
    }

    return [null, $_scount];
}