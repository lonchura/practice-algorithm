<?php
//带顺序的全排列问题
//假定有4组8个节点((A0,A1),(B0,B1),(C0,C1),(D0,D1))
//合法条件是A1必须在A0之后, B1必须在B0之后
//求打印所有可行的全排列

$list = ['A0','A1','B0','B1','C0','C1','D0','D1'];
//$list = ['A0','A1', 'B0'];
$inst = new PermutationSort($list);
//var_dump(memory_get_usage(true));
$inst->execute();
//var_dump(memory_get_usage(true));

class PermutationSort {

    private $list;

    /**
     * Permutation constructor.
     * @param array $list
     */
    public function __construct(array $list)
    {
        $this->list = $list;
    }

    /**
     * execute
     */
    public function execute() {
        $result = $this->qpSortList($this->list);
        foreach($result as $list) {
            // 检测条件
            $rList = array_combine(array_values($list), array_keys($list));
            if($rList['A0'] > $rList['A1']
                || $rList['B0'] > $rList['B1'])
            {
                continue ;
            }

            // 输出
            foreach ($list as $v) {
                echo "$v\t";
            }
            echo "\n";
        }
    }

    /**
     * @param array $listPart
     * @return array
     *  (
     *      array (...),
     *      ...
     *  )
     */
    private function qpSortList(array $listPart=[]) {
        if(count($listPart) == 1) {
            return array($listPart);
        }

        $result = array();
        foreach ($listPart as $k=>$val) {
            // 剩余list
            $nextPartList = $listPart;
            unset($nextPartList[$k]);

            // 剩余list全排序
            $list = $this->qpSortList($nextPartList);
            foreach($list as $partList) {
                $result[] = array_merge([$val], $partList);
            }
        }

        return $result;
    }

}