<?php
// 源数组
$list = [1987, 100, 9, 250, 9, 1024, 925];
// 排序后数组
$sorted = insertSort($list);
// 打印数组
foreach($sorted as $int) {
    printf("%d\n", $int);
}

/**
 * bubble sort
 * @param array $list
 * @return array
 */
function insertSort(array $list=[]) {
    // 获取数量
    $count = count($list);

    // 插入次数n-1次
    for($i=1; $i<=$count-1; $i++) {
        // - 首次  至多对比1次 从2个开始往前对比(10)
        // - 第二次  至多对比2次 从第3个开始往前对比(21 10)
        // - 第i次  至多对比i次 从i+1个开始往前对比(32 21 10)
        for($j=$i; $j>=1; $j--) {
            // 从后往前对比，如果后者小于前者，就将小的往前迁移（数组索引从0开始）
            if($list[$j] < $list[$j-1]) {
                $_tmp = $list[$j];
                $list[$j] = $list[$j-1];
                $list[$j-1] = $_tmp;
            }
        }
    }

    return $list;
}