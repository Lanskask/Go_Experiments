
def tupleF[F[_], A, B] (els1: F[A], el2: F[B]) (implicit: F[(A, B)]) = {

}

trait TupleF[F[_]] {
  def unit[F[_], A, B](el1: F[A], el2: F[B]): F[(A, B)]
}

implicit object TupleFList extends TupleF[List[_]] {
  def unit[A, B](el1: List[A], el2: List[B]): List[(A, B)] = {
    el1.zip(el2)
  }
}

val a = tupleF(List(1), List(2))
val b = tupleF(Some('A'), Some('B'))
