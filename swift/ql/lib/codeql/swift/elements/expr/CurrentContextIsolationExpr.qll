// generated by codegen/codegen.py, do not edit
/**
 * This module provides the public class `CurrentContextIsolationExpr`.
 */

private import internal.CurrentContextIsolationExprImpl
import codeql.swift.elements.expr.Expr

/**
 * An expression that extracts the actor isolation of the current context, of type `(any Actor)?`.
 * This is synthesized by the type checker and does not have any way to be expressed explicitly in
 * the source.
 */
final class CurrentContextIsolationExpr = Impl::CurrentContextIsolationExpr;
