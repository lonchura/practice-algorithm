<?php
// 源数组
$list = [1987, 100, 9, 250, 9, 1024, 925];
// 排序后数组
$sorted = selectSort($list);
foreach($sorted as $int) {
    printf("%d\n", $int);
}

/**
 * 选择排序算法
 * @param array $list
 * @return array
 */
function selectSort(array $list) {
    // check empty
    if (empty($list)) {
        return $list;
    }

    // init sort list
    $sortedList = [];

    // sort
    $sortCount = count($list);
    while($sortCount-- > 0) {
        $smallestIndex = findSmallest($list);
        // add to sorted list
        $sortedList[] = $list[$smallestIndex];
        // remove smallest in list
        unset($list[$smallestIndex]);
        // reset list index
        $list = array_values($list);
    }

    return $sortedList;
}

/**
 * find smallest index of list
 * @param array $list
 * @return int smallest index
 */
function findSmallest(array $list) {
    $smallTmpIndex = 0;
    $smallTmp = $list[0];

    // find smallest
    $count = count($list);
    for($i=1; $i<$count; $i++) {
        if($smallTmp > $list[$i]) {
            $smallTmpIndex = $i;
            $smallTmp = $list[$i];
        }
    }

    return $smallTmpIndex;
}